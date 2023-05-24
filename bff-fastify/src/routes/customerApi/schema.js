exports.getAllCustomers = {
    description: 'Anyone can access to it',
    tags: ['customer-service'],
    summary: 'Get all customers',
    querystring: {
        limit: { type: 'number' },
        offset: { type: 'number' },
        orderDirection: { type: 'string', description: 'ASC or DESC' },
        orderBy: { type: 'string', description: 'name' },
        isCount: { type: 'boolean', default: false }
    },
    response: {
        // 200: {
        //     description: 'Successful response',
        //     type: 'object',
        //     properties: {
        //         items: {
        //             type: 'array',
        //             items: {
        //                 type: 'object',
        //                 properties: {
        //                     name: { type: 'string' },
        //                     followerChangesCount1Day: { type: 'number' },
        //                     followerChangesCount7Day: { type: 'number' },
        //                     followerChangesCount30Day: { type: 'number' },
        //                     visitCount: { type: 'number' },
        //                     orderCount: { type: 'number' },
        //                     priceValue: { type: 'number' }
        //                 }
        //             }
        //         },
        //         totalCount: { type: 'number', example: '4' },
        //         limit: { type: 'number', example: '25' },
        //         offset: { type: 'number', example: '0' }
        //     }
        // },
        500: {
            description: 'Internal Server Error',
            type: 'object',
            properties: {
                message: { type: 'string' },
                correlationId: { type: 'string' },
                statusCode: { type: 'number' },
                errorCode: { type: 'number' }
            }
        }
    }
}

exports.getCustomer = {
    description: 'Anyone can access to it',
    tags: ['customer-service'],
    summary: 'Get customer by id',
    path: '/api/v1/customers/{customerId}',
    params: {
        type: 'object',
        required: ['customerId'],
        properties: {
            customerId: {
                type: 'string'
            }
        }
    },
    response: {
        500: {
            description: 'Internal Server Error',
            type: 'object',
            properties: {
                message: { type: 'string' },
                correlationId: { type: 'string' },
                statusCode: { type: 'number' },
                errorCode: { type: 'number' }
            }
        }
    }
}

exports.getCustomerValidation = {
    description: 'Anyone can access to it',
    tags: ['customer-service'],
    summary: 'Get customer validation',
    path: '/api/v1/customers/validation/{customerId}',
    params: {
        type: 'object',
        required: ['customerId'],
        properties: {
            customerId: {
                type: 'string'
            }
        }
    },
    response: {
        500: {
            description: 'Internal Server Error',
            type: 'object',
            properties: {
                message: { type: 'string' },
                correlationId: { type: 'string' },
                statusCode: { type: 'number' },
                errorCode: { type: 'number' }
            }
        }
    }
}

exports.postCreateCustomer = {
    description: 'Anyone can access to it',
    tags: ['customer-service'],
    summary: 'Create a customer',
    body: {
        type: 'object',
        properties: {
            name: { type: 'string' },
            email: { type: 'string' },
            address: {
                type: 'object',
                properties: {
                    addressLine: { type: 'string' },
                    city: { type: 'string' },
                    country: { type: 'string' },
                    cityCode: { type: 'number' }
                }
            }
        }
    },
    response: {
        201: {
            description: 'Success response',
            type: 'object',
            properties: {
            }
        },
        400: {
            description: 'Bad Request',
            type: 'object',
            properties: {
                message: { type: 'string' },
                statusCode: { type: 'number' },
                errorCode: { type: 'number' }
            }
        },
        500: {
            description: 'Internal Server Error',
            type: 'object',
            properties: {
                message: { type: 'string' },
                statusCode: { type: 'number' },
                errorCode: { type: 'number' }
            }
        }
    }
}

exports.postUpdateCustomer = {
    description: 'Anyone can access to it',
    tags: ['customer-service'],
    summary: 'Update a customer',
    body: {
        type: 'object',
        properties: {
            id: { type: 'string' },
            name: { type: 'string' },
            email: { type: 'string' },
            address: {
                type: 'object',
                properties: {
                    addressLine: { type: 'string' },
                    city: { type: 'string' },
                    country: { type: 'string' },
                    cityCode: { type: 'number' }
                }
            }
        }
    },
    response: {
        201: {
            description: 'Success response',
            type: 'object',
            properties: {
            }
        },
        400: {
            description: 'Bad Request',
            type: 'object',
            properties: {
                message: { type: 'string' },
                statusCode: { type: 'number' },
                errorCode: { type: 'number' }
            }
        },
        500: {
            description: 'Internal Server Error',
            type: 'object',
            properties: {
                message: { type: 'string' },
                statusCode: { type: 'number' },
                errorCode: { type: 'number' }
            }
        }
    }
}

exports.deleteCustomer = {
    description: 'Anyone can access to it',
    tags: ['customer-service'],
    summary: 'Delete customer by id',
    path: '/api/v1/customers/{customerId}',
    params: {
        type: 'object',
        required: ['customerId'],
        properties: {
            customerId: {
                type: 'string'
            }
        }
    },
    response: {
        500: {
            description: 'Internal Server Error',
            type: 'object',
            properties: {
                message: { type: 'string' },
                correlationId: { type: 'string' },
                statusCode: { type: 'number' },
                errorCode: { type: 'number' }
            }
        }
    }
}
