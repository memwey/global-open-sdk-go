# Change Log

## Tax Refund API Support

新增 Alipay+ 退税（Tax Refund）相关接口的 SDK 支持，复用现有的签名、校验、加密、错误处理等基础能力。

### 新增文件

#### Model

| 文件 | 说明 |
|------|------|
| `com/alipay/api/model/Passport.go` | 护照信息模型，包含 `fullName`、`passportNumber`、`nationality`、`valideDate`、`expireDate`、`birthday` |
| `com/alipay/api/model/TaxRefundUser.go` | 退税用户模型，包含 `userId`、`userLoginId`、`userName` |
| `com/alipay/api/model/TaxRefundQuote.go` | 退税汇率模型，与通用 `Quote` 的区别在于 `quotePrice` 为 `string` 类型（API 返回字符串格式） |
| `com/alipay/api/model/TaxRefundFormStatusType.go` | 退税表单状态枚举，包含 `INIT`、`STAMPED`、`REJECTED_BY_CUSTOMS`、`RECEIVED`、`VOIDED`、`FAILED`、`EXPIRED`、`REFUNDED`、`REFUNDED_NON_ALIPAYPLUS` |

#### Request

| 文件 | 接口 | API Path |
|------|------|----------|
| `com/alipay/api/request/taxrefund/AlipayEvaluateOriginalCreditRequest.go` | 评估 OCT 是否可用 | `POST /aps/api/v1/funds/evaluateOriginalCredit` |
| `com/alipay/api/request/taxrefund/AlipayCreateOriginalCreditRequest.go` | 创建 OCT（发起退税） | `POST /aps/api/v1/funds/createOriginalCredit` |
| `com/alipay/api/request/taxrefund/AlipayInquireOriginalCreditRequest.go` | 查询 OCT 结果 | `POST /aps/api/v1/funds/inquireOriginalCredit` |
| `com/alipay/api/request/taxrefund/AlipayConfirmOriginalCreditRequest.go` | 确认 OCT 结果 | `POST /aps/api/v1/funds/confirmOriginalCredit` |
| `com/alipay/api/request/taxrefund/AlipaySyncTaxRefundFormRequest.go` | 同步退税表单状态 | `POST /aps/api/v1/funds/syncTaxRefundForm` |

#### Response

| 文件 | 说明 |
|------|------|
| `com/alipay/api/response/taxrefund/AlipayEvaluateOriginalCreditResponse.go` | 包含 `acquirerId`、`pspId`、`payeeAmount`、`payeeQuote`、`payee`、`passport`、`walletBrandName` |
| `com/alipay/api/response/taxrefund/AlipayCreateOriginalCreditResponse.go` | 包含 `originalCreditId`、`originalCreditTime`、`payeeAmount`、`payeeQuote`、`payee` |
| `com/alipay/api/response/taxrefund/AlipayInquireOriginalCreditResponse.go` | 包含 `originalCreditResult`（OCT 本身结果）、`payerAmount`、`payeeAmount`、`payer`、`payee` 等完整字段 |
| `com/alipay/api/response/taxrefund/AlipayConfirmOriginalCreditResponse.go` | 包含 `result`、`acquirerId`、`pspId` |
| `com/alipay/api/response/taxrefund/AlipaySyncTaxRefundFormResponse.go` | 包含 `result` |

#### Notify

| 文件 | 说明 |
|------|------|
| `com/alipay/api/request/notify/AlipayOriginalCreditResultNotify.go` | 接收 Alipay+ 主动推送的 OCT 结果通知（Alipay+ 作为调用方，我方作为响应方） |

#### Example

| 文件 | 说明 |
|------|------|
| `com/alipay/example/taxrefund_demo.go` | 五个接口的完整使用示例：`evaluateOriginalCredit`、`createOriginalCredit`、`inquireOriginalCredit`、`confirmOriginalCredit`、`syncTaxRefundForm` |

---

### 接口说明

#### evaluateOriginalCredit

调用前置评估接口，判断用户 MPP 账户是否支持 OCT（退税转账）。

必填字段：`payerAmount`、`payer`、`payeeMethod`、`evaluationType`（`BY_CODE` 或 `BY_USER_ID`）、`scenarioType`、`subScenarioType`、`departureRegion`

#### createOriginalCredit

正式发起退税 OCT。`scenarioType=TAX_REFUND` 时，`taxRefundFormNumber`、`departureRegion`、`departurePort`、`totalSalesAmount` 为必填。

必填字段：`originalCreditRequestId`、`payerAmount`、`payer`、`payee`、`scenarioType`、`subScenarioType`

#### inquireOriginalCredit

查询已创建的 OCT 状态。`originalCreditRequestId` 与 `originalCreditId` 二选一传入。

响应中 `originalCreditResult` 为 OCT 本身的处理结果，与外层 `result`（查询请求本身的结果）相互独立。

#### notifyOriginalCredit（Notify）

Alipay+ 在 OCT 达到终态后，主动推送到 `createOriginalCredit` 请求中指定的 `payerNotificationUrl`。

接收方处理流程：
1. 使用 `tools.CheckSignature` 验证请求签名
2. 将请求体反序列化为 `AlipayOriginalCreditResultNotify`
3. 处理业务逻辑
4. 返回 `{"result": {"resultStatus": "S", "resultCode": "SUCCESS", "resultMessage": "Success"}}`

未返回成功时，Alipay+ 将按 2min / 10min / 10min / 1h / 2h / 6h / 15h 间隔最多重试 7 次。

#### confirmOriginalCredit

当 `inquireOriginalCredit` 查询 1 分钟后 OCT 状态仍未知时，调用此接口向 Alipay+ 确认 OCT 是否成功。

必填字段：`originalCreditRequestId` 与 `originalCreditId` 二选一

响应中 `result.resultStatus=S` 表示确认成功（OCT 成功），`result.resultStatus=F` 表示确认失败（OCT 失败），需要根据 `result.resultCode` 区分处理。

#### syncTaxRefundForm

同步退税表单状态给 Alipay+。

必填字段：`taxRefundFormNumber`、`formStatus`、`taxRefundAmount`、`merchants`、`userId`

`formStatus` 可选值：`INIT`、`STAMPED`、`REJECTED_BY_CUSTOMS`、`RECEIVED`、`VOIDED`、`FAILED`、`EXPIRED`、`REFUNDED`、`REFUNDED_NON_ALIPAYPLUS`

---

### 注意事项

- 签名、验签、请求头结构与现有接口完全一致，复用 `tools.GenSign` / `tools.CheckSignature`，无需额外配置。
- 税退接口路径前缀为 `/aps/api/v1/funds/...`，与支付接口的 `/ams/api/v1/...` 不同。当前 `AdjustSandboxUrl` 仅处理 `/ams/api` 前缀的沙箱替换，**沙箱环境下税退接口的路径替换逻辑尚未覆盖**，如需使用沙箱请注意。
- `Passport.valideDate` 字段名来自 API sample response 原文，如后续联调发现实际字段名为 `validDate` 需同步修正。
