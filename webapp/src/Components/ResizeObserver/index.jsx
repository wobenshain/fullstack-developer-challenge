import React, {Component, createRef} from "react";
import ResizeObserver from "resize-observer-polyfill";
import {notification} from "antd";
import { indexOutOfBounds } from 'Consts/errorCodes';

// I wrote this to do a thing but realized I didn't need it
// Keeping it in here because it's a useful bit of code
class ResizeChecker extends Component {
    constructor(props) {
        super(props);

        this.viewport = createRef();

        this.state = {
            observer: null,
        };
    }

    componentDidMount() {
        const { onChange } = this.props;
        const observer = new ResizeObserver(observed => {
            if (observed.length !== 1) {
                notification.warn(indexOutOfBounds);
                if (observed.length < 1) return;
            }
            const { target } = observed[0];
            onChange({
                height: target.offsetHeight,
                width: target.offsetWidth,
            });
        });
        this.setState({ observer }, () => {
            observer.observe(this.viewport.current);
        });
    }

    render() {
        const { children } = this.props;
        return <span ref={this.viewport}>{children}</span>;
    }
}

export default ResizeChecker;
