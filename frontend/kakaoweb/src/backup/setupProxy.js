const proxy = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(proxy('/api', { target: 'http://localhost:10102/' }));
  app.use(proxy('/*.svg', { target: 'http://localhost:10102/' }));
};
