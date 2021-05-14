// module.exports = {
//   devServer: {
//     proxy: {
//       "^/api": {
//         target: "http://localhost:5555",
//         ws: true,
//         changeOrigin: true,
//         pathRewrite: { "^/api": "/" }
//       },
//     },
//   },
// }

module.exports = {
  devServer: {
    proxy: {
      "": {
        target: "http://localhost:5555",
        ws: true,
        changeOrigin: true,
        pathRewrite: { "^/api": "/" }
      },
    },
  },
}

