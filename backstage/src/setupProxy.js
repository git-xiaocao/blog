const {createProxyMiddleware} = require("http-proxy-middleware");

module.exports = (app) => {
    app.use(
        "/api",
        createProxyMiddleware(
            {
                target: "http://localhost:44444/api/backstage",
                changeOrigin: true,
                pathRewrite: {"^/api": ""},
            })
    );
};