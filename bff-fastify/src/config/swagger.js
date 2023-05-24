exports.options = {
    routePrefix: '/swagger',
    exposeRoute: true,
    openapi: {
        info: {
            title: 'bff',
            description: 'Backend for frontend layer of Order Application',
            version: '1.0.1'
        },
        host: 'localhost',
        consumes: ['application/json'],
        produces: ['application/json']
        // components: {
        //     securitySchemes: {
        //         Bearer: {
        //             type: 'http',
        //             scheme: 'bearer',
        //             bearerFormat: 'JWT'
        //         }
        //     }
        // },
        // security: [{
        //     Bearer: []
        // }]
    }
}
