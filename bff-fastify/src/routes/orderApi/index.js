const orderController = require('../../controllers/orderApi/index')
const schema = require('./schema')
const customerController = require("../../controllers/customerApi");

const orderRoutes = [
    {
        method: 'GET',
        url: '/api/v1/orders',
        handler: orderController.getAllOrders,
        schema: schema.getAllOrders
    },
    {
        method: 'GET',
        url: '/api/v1/orders/order/:orderId',
        handler: orderController.getOrderByOrderId,
        schema: schema.getOrderByOrderId
    },
    {
        method: 'GET',
        url: '/api/v1/orders/:customerId',
        handler: orderController.getOrdersByCustomerId,
        schema: schema.getOrdersByCustomerId
    },
    {
        method: 'POST',
        url: '/api/v1/orders/create',
        handler: orderController.postCreateOrder,
        schema: schema.postCreateOrder
    },
    {
        method: 'POST',
        url: '/api/v1/orders/update',
        handler: orderController.postUpdateOrder,
        schema: schema.postUpdateOrder
    },
    {
        method: 'POST',
        url: '/api/v1/orders/change-order-status',
        handler: orderController.postUpdateOrderStatus,
        schema: schema.postUpdateOrderStatus
    },
    {
        method: 'DELETE',
        url: '/api/v1/orders/:orderId',
        handler: orderController.deleteOrder,
        schema: schema.deleteOrder
    }
]

module.exports = orderRoutes
