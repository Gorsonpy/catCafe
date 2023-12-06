namespace go order

struct BaseRequest{
}

struct BaseResponse{
    1: i64 code,
    2: string msg
}

struct QueryOrdersReq{
    1:bool status
    2:i64 limit
}

struct OrderModel{
    1:i64 orderId
    2:i64 customerId
    3:i64 catId
    4:string orderDetails
    5:double totalAmount
    6:bool paymentStatus
    7:string paymentMethod
    8:string orderTime 
}
struct QueryOrderResp{
    1: BaseResponse base
    2: list<OrderModel> data
}

struct CreateOrderReq{
    1: string orderDetails
    2: double totalAmount
    3: string paymentMethod
    4: i64 catId
}

struct CreateOrderResp{
    1: BaseResponse base
    2: map<string, i64> data
}

struct ConfirmOrderReq{
    1: i64 orderId
    2: bool status
}
struct UpdateOrderReq{
    1: i64 orderId
    2: string orderDetails
    3: double totalAmount
}

struct DelOrderReq{
    1: list<i64> orderId
}

service OrderService{
    QueryOrderResp queryOrders(1:QueryOrdersReq req)(api.post = "/order/query")
    CreateOrderResp createOrder(1:CreateOrderReq req)(api.post = "/order")
    BaseResponse confirmOrder(1:ConfirmOrderReq req)(api.put = "/order/confirm")
    BaseResponse UpdateOrder(1:UpdateOrderReq req)(api.put = "/order")
    BaseResponse DelOrder(1:DelOrderReq req)(api.delete = "/order")
}

