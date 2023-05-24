exports.getAllCustomers = {
    description: 'Anyone can access to it',
    tags: ['customer'],
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
