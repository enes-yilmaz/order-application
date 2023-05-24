const axios = require('axios')
const config = require('../config/index')
const path = require('path')

function _get (args){
    const instance = newInstance(args)
    const data = args.query ? { params: args.query } : undefined

    return instance.get(args.url, data).catch((error) => {
        console.log(JSON.stringify({
            statusCode: (error.response || {}).status,
            message: ('Path: ' + args.baseURL || '') + args.url,
            errorCode: 9999
        }))
        return error.response
    })
}

function _post (args) {
    const instance = newInstance(args)
    return instance.post(args.url, args.body).catch((error) => {
        console.log(JSON.stringify({
            statusCode: (error.response || {}).status,
            message: ('Path: ' + args.baseURL || '') + args.url,
            errorCode: 9999
        }))
        return error.response
    })
}

function _put (args) {
    const instance = newInstance(args)
    return instance.put(args.url, args.body).catch((error) => {
        console.log(JSON.stringify({
            statusCode: (error.response || {}).status,
            message: ('Path: ' + args.baseURL || '') + args.url,
            errorCode: 9999
        }))
        return error.response
    })
}

function _delete (args) {
    const instance = newInstance(args)
    return instance.delete(args.url, args.body).catch((error) => {
        console.log(JSON.stringify({
            statusCode: (error.response || {}).status,
            message: ('Path: ' + args.baseURL || '') + args.url,
            errorCode: 9999
        }))
        return error.response
    })
}

function newInstance({
                         url,
                         baseURL,
                         body = {},
                         options = {},
                         cacheOptions,
                         token
                     }){
    let headers = {
        'Content-Type': 'application/json',
        Accept: 'application/json'
    }
    // if (token) {
    //     headers.Authorization = token
    // }
    if (options.headers) {
        headers = { ...headers, ...options.headers }
    }
    const i = axios.create({
        baseURL,
        headers,
        timeout: options.timeout || 100000,
        maxContentLength: 100000000,
        maxBodyLength: 1000000000
    })
    i.customOptions = options
    return i
}

module.exports = {
    async get (args) {
        return _get(args)
    },
    async post (args) {
        return _post(args)
    },
    async put (args) {
        return _put(args)
    },
    async delete (args) {
        return _delete(args)
    }
}
