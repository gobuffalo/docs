'use strict';

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }(); /* eslint no-console: 0 */

var _fs = require('fs');

var _fs2 = _interopRequireDefault(_fs);

var _lodash = require('lodash.union');

var _lodash2 = _interopRequireDefault(_lodash);

var _getFiles = require('./getFiles');

var _getFiles2 = _interopRequireDefault(_getFiles);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

var WebpackCleanupPlugin = function () {
  function WebpackCleanupPlugin() {
    var options = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};

    _classCallCheck(this, WebpackCleanupPlugin);

    this.options = options;
  }

  _createClass(WebpackCleanupPlugin, [{
    key: 'apply',
    value: function apply(compiler) {
      var _this = this;

      var outputPath = compiler.options.output.path;

      compiler.plugin('done', function (stats) {
        if (compiler.outputFileSystem.constructor.name !== 'NodeOutputFileSystem') {
          return;
        }

        var assets = stats.toJson().assets.map(function (asset) {
          return asset.name;
        });
        var exclude = (0, _lodash2.default)(_this.options.exclude, assets);
        var files = (0, _getFiles2.default)(outputPath, exclude);

        if (_this.options.preview) {
          console.log('%s file(s) would be deleted:', files.length);
          files.forEach(function (file) {
            return console.log('    %s', file);
          });
          console.log();
        } else {
          files.forEach(_fs2.default.unlinkSync);
        }
        if (!_this.options.quiet) {
          console.log('\nWebpackCleanupPlugin: %s file(s) deleted.', files.length);
        }
      });
    }
  }]);

  return WebpackCleanupPlugin;
}();

module.exports = WebpackCleanupPlugin;
//# sourceMappingURL=WebpackCleanupPlugin.js.map