
var ElapserExtGlobal = React.createClass({
    render: function () {
        var message =
            'Global JavaScript timer has been successfully running for ' + this.props.elapsed + ' seconds.';

        return React.DOM.div(null, message)
    },

    shouldComponentUpdate: function (nextProps, nextState) {
        return  nextProps.elapsed !== this.props.elapsed;
    }

});
