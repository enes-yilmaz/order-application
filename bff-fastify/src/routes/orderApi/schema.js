exports.getAllOrders = {
    description: 'Anyone can access to it',
    tags: ['order-service'],
    summary: 'Get all orders',
    querystring: {
        limit: { type: 'number' },
        offset: { type: 'number' },
        orderDirection: { type: 'string', description: 'ASC or DESC' },
        orderBy: { type: 'string', description: 'name' },
        isCount: { type: 'boolean', default: false }
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

exports.getOrderByOrderId = {
    description: 'Anyone can access to it',
    tags: ['order-service'],
    summary: 'Get order by orderId',
    path: '/api/v1/orders/order/{orderId}',
    params: {
        type: 'object',
        required: ['orderId'],
        properties: {
            orderId: {
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

exports.getOrdersByCustomerId = {
    description: 'Anyone can access to it',
    tags: ['order-service'],
    summary: 'Get orders by customerId',
    path: '/api/v1/orders/{customerId}',
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

exports.postCreateOrder = {
    description: 'Anyone can access to it',
    tags: ['order-service'],
    summary: 'Create an order',
    body: {
        type: 'object',
        properties: {
            customerId: { type: 'string' },
            quantity: { type: 'number' },
            price: { type: 'number' },
            status: { type: 'string' },
            address: {
                type: 'object',
                properties: {
                    addressLine: { type: 'string' },
                    city: { type: 'string' },
                    country: { type: 'string' },
                    cityCode: { type: 'number' }
                }
            },
            product: {
                type: 'object',
                properties: {
                    id: { type: 'string' },
                    imageUrl: { type: 'string' },
                    name: { type: 'string' }
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

exports.postUpdateOrder = {
    description: 'Anyone can access to it',
    tags: ['order-service'],
    summary: 'Update an order',
    body: {
        type: 'object',
        properties: {
            orderId: { type: 'string' },
            customerId: { type: 'string' },
            quantity: { type: 'number' },
            price: { type: 'number' },
            status: { type: 'string' },
            address: {
                type: 'object',
                properties: {
                    addressLine: { type: 'string' },
                    city: { type: 'string' },
                    country: { type: 'string' },
                    cityCode: { type: 'number' }
                }
            },
            product: {
                type: 'object',
                properties: {
                    id: { type: 'string' },
                    imageUrl: { type: 'string' },
                    name: { type: 'string' }
                }
            }
        }
    },
    response: {
        200: {
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

exports.postUpdateOrderStatus = {
    description: 'Anyone can access to it',
    tags: ['order-service'],
    summary: 'Update an order status',
    body: {
        type: 'object',
        properties: {
            orderId: { type: 'string' },
            status: { type: 'string' }
        }
    },
    response: {
        204: {
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

exports.deleteOrder = {
    description: 'Anyone can access to it',
    tags: ['order-service'],
    summary: 'Delete order',
    path: '/api/v1/orders/{orderId}',
    params: {
        type: 'object',
        required: ['orderId'],
        properties: {
            orderId: {
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
