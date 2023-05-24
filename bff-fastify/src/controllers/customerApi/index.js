const services = require('../../services/index')
const helpers = require('./helpers')
const { isEmptyObject } = require('../../utils')

exports.getAllCustomers = async (req, res) => {
    let {
        limit,
        offset,
        orderDirection,
        orderBy,
        isCount
    } = req.query

    limit = limit < 100 && limit > 0 ? +limit : 25
    offset = offset >= 0 ? +offset : 0

    const response = await services.customerApi.getAllCustomers({ limit, offset, orderDirection, orderBy, isCount })
    return res.status(response?.status).send(response?.data)
}

exports.getCustomer = async (req, res) => {
    const customerId = req?.params.customerId
    const response = await services.customerApi.getCustomer(customerId)
    return res.status(response?.status).send(response?.data)
}

exports.getCustomerValidation = async (req, res) => {
    const customerId = req?.params.customerId
    const response = await services.customerApi.getCustomerValidation(customerId)
    return res.status(response?.status).send(response?.data)
}

exports.postCreateCustomer = async (req, res) => {
    const body = req?.body
    if (isEmptyObject(body?.data)) {
        body.data = {}
    }

    const response = await services.customerApi.postCreateCustomer(body)
    return res.status(response?.status).send(response?.data)
}

exports.postUpdateCustomer = async (req, res) => {
    const body = req?.body
    if (isEmptyObject(body?.data)) {
        body.data = {}
    }

    const response = await services.customerApi.postUpdateCustomer(body)
    return res.status(response?.status).send(response?.data)
}

exports.deleteCustomer = async (req, res) => {
    const customerId = req?.params.customerId
    const response = await services.customerApi.deleteCustomer(customerId)
    return res.status(response?.status).send(response?.data)
}
