const client = require('../../utils/httpHelper')
const config = require('../../config')
const CustomError = require('../../errors')
const { isEmptyObject } = require('../../utils')

exports.getAllCustomers = async ({ limit, offset, orderDirection, orderBy, isCount }) => {
    const query = { limit, offset, orderDirection, orderBy, isCount }
    // const url = '/customers?' + Object.keys(query).filter(k => !!query[k]).map(k => `${k}=${encodeURIComponent(query[k])}`).join('&')
    const url = '/customers?' + Object.keys(query).filter(k => !!query[k]).map(k => `${k}=${query[k]}`).join('&')

    const response = await client.get({
        baseURL : config.services.internals.customerApi,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Customer service failed to get all customers!',
            errorCode: 10101
        })
    } else if (response?.status === 400) {
        throw new CustomError({
            statusCode: 400,
            message: response?.data.Desc,
            errorCode: 10102
        })
    }
    return response
}

exports.getCustomer = async (customerId) => {
    const url =`/customers/${customerId || ''}`
    const response = await client.get({
        baseURL : config.services.internals.customerApi,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Customer service failed to get customer by customerId!',
            errorCode: 10103
        })
    }
    return response
}

exports.getCustomerValidation = async (customerId) => {
    const url =`/customers/validate/${customerId || ''}`
    const response = await client.get({
        baseURL : config.services.internals.customerApi,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Customer service failed to get customer validation!',
            errorCode: 10104
        })
    }
    return response
}

exports.postCreateCustomer = async (body) => {
    const url ='/customers'
    const response = await client.post({
        baseURL : config.services.internals.customerApi,
        body: body,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Customer service failed to create customer!',
            errorCode: 10105
        })
    } else if (response?.status === 400) {
        throw new CustomError({
            statusCode: 400,
            message: response?.data.Desc,
            errorCode: 10106
        })
    }
    return response
}

exports.postUpdateCustomer = async (body) => {
    const url ='/customers/update'
    const response = await client.post({
        baseURL : config.services.internals.customerApi,
        body: body,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Customer service failed to update customer!',
            errorCode: 10107
        })
    } else if (response?.status === 400) {
        throw new CustomError({
            statusCode: 400,
            message: response?.data.Desc,
            errorCode: 10108
        })
    }
    return response
}

exports.deleteCustomer = async (customerId) => {
    const url =`/customers/${customerId || ''}`
    const response = await client.delete({
        baseURL : config.services.internals.customerApi,
        url
    })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Customer service failed to delete customer by customerId!',
            errorCode: 10109
        })
    }
    return response
}
