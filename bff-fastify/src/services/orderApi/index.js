const clients = require('../../clients')

exports.getAllOrders = async ({ limit, offset, orderDirection, orderBy, isCount }) => {
    return await clients.orderApi.getAllOrders({ limit, offset, orderDirection, orderBy, isCount })
}

exports.getOrderByOrderId = async (orderId) => {
    return await clients.orderApi.getOrderByOrderId(orderId)
}

exports.getOrdersByCustomerId = async (customerId) => {
    return await clients.orderApi.getOrdersByCustomerId(customerId)
}

exports.postCreateOrder = async (body) => {
    return await clients.orderApi.postCreateOrder(body)
}

exports.postUpdateOrder = async (body) => {
    return await clients.orderApi.postUpdateOrder(body)
}

exports.postUpdateOrderStatus = async (body) => {
    return await clients.orderApi.postUpdateOrderStatus(body)
}

exports.deleteOrder = async (orderId) => {
    return await clients.orderApi.deleteOrder(orderId)
}
