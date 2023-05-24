const services = require('../../services/index')
const { isEmptyObject } = require('../../utils')

exports.getAllOrders = async (req, res) => {
    let {
        limit,
        offset,
        orderDirection,
        orderBy,
        isCount
    } = req.query

    limit = limit < 100 && limit > 0 ? +limit : 25
    offset = offset >= 0 ? +offset : 0

    const response = await services.orderApi.getAllOrders({ limit, offset, orderDirection, orderBy, isCount })
    return res.status(response?.status).send(response?.data)
}

exports.getOrderByOrderId = async (req, res) => {
    const orderId = req?.params.orderId
    const response = await services.orderApi.getOrderByOrderId(orderId)
    return res.status(response?.status).send(response?.data)
}

exports.getOrdersByCustomerId = async (req, res) => {
    const customerId = req?.params.customerId
    const response = await services.orderApi.getOrdersByCustomerId(customerId)
    return res.status(response?.status).send(response?.data)
}

exports.postCreateOrder = async (req, res) => {
    const body = req?.body
    if (isEmptyObject(body?.data)) {
        body.data = {}
    }

    const response = await services.orderApi.postCreateOrder(body)
    return res.status(response?.status).send(response?.data)
}

exports.postUpdateOrder = async (req, res) => {
    const body = req?.body
    if (isEmptyObject(body?.data)) {
        body.data = {}
    }

    const response = await services.orderApi.postUpdateOrder(body)
    return res.status(response?.status).send(response?.data)
}

exports.postUpdateOrderStatus = async (req, res) => {
    const body = req?.body
    if (isEmptyObject(body?.data)) {
        body.data = {}
    }

    const response = await services.orderApi.postUpdateOrderStatus(body)
    return res.status(response?.status).send(response?.data)
}

exports.deleteOrder = async (req, res) => {
    const orderId = req?.params.orderId
    const response = await services.orderApi.deleteOrder(orderId)
    return res.status(response?.status).send(response?.data)
}
