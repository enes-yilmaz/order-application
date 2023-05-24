const client = require('../../utils/httpHelper')
const config = require('../../config')
const CustomError = require('../../errors')
const { isEmptyObject } = require('../../utils')

exports.getAllOrders = async ({ limit, offset, orderDirection, orderBy, isCount }) => {
    const query = { limit, offset, orderDirection, orderBy, isCount }
    // const url = '/customers?' + Object.keys(query).filter(k => !!query[k]).map(k => `${k}=${encodeURIComponent(query[k])}`).join('&')
    const url = '/orders?' + Object.keys(query).filter(k => !!query[k]).map(k => `${k}=${query[k]}`).join('&')

    const response = await client.get({
        baseURL : config.services.internals.orderApi,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Order service failed to get all orders!',
            errorCode: 10111
        })
    } else if (response?.status === 400) {
        throw new CustomError({
            statusCode: 400,
            message: response?.data.Desc,
            errorCode: 10112
        })
    }
    return response
}

exports.getOrderByOrderId = async (orderId) => {
    const url =`/orders/order/${orderId || ''}`
    const response = await client.get({
        baseURL : config.services.internals.orderApi,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Order service failed to get customer by orderId!',
            errorCode: 1013
        })
    }
    return response
}

exports.getOrdersByCustomerId = async (customerId) => {
    const url =`/orders/${customerId || ''}`
    const response = await client.get({
        baseURL : config.services.internals.orderApi,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Order service failed to get customer by orderId!',
            errorCode: 10114
        })
    }
    return response
}

exports.postCreateOrder = async (body) => {
    const url ='/orders/create'
    const response = await client.post({
        baseURL : config.services.internals.orderApi,
        body: body,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Order service failed to create order!',
            errorCode: 10115
        })
    } else if (response?.status === 400) {
        throw new CustomError({
            statusCode: 400,
            message: response?.data.Desc,
            errorCode: 10116
        })
    }
    return response
}

exports.postUpdateOrder = async (body) => {
    const url ='/orders/update'
    const response = await client.post({
        baseURL : config.services.internals.orderApi,
        body: body,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Order service failed to update order!',
            errorCode: 10117
        })
    } else if (response?.status === 400) {
        throw new CustomError({
            statusCode: 400,
            message: response?.data.Desc,
            errorCode: 10118
        })
    }
    return response
}

exports.postUpdateOrderStatus = async (body) => {
    const url ='/orders/change-order-status'
    const response = await client.post({
        baseURL : config.services.internals.orderApi,
        body: body,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Order service failed to update order status!',
            errorCode: 10119
        })
    } else if (response?.status === 400) {
        throw new CustomError({
            statusCode: 400,
            message: response?.data.Desc,
            errorCode: 10120
        })
    }
    return response
}

exports.deleteOrder = async (orderId) => {
    const url =`/orders/${orderId || ''}`
    const response = await client.delete({
        baseURL : config.services.internals.orderApi,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Order service failed to delete order!',
            errorCode: 10121
        })
    }
    return response
}
