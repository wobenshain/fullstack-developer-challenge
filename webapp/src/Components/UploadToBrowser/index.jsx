import React, { Component, createRef /*, useEffect, useState*/ } from 'react';

class Wrapped extends Component {
    constructor(props) {
        super(props);

        this.viewport = createRef()
    }

    render() {
        const { children } = this.props;
        return <div ref={this.viewport}>{children}</div>;
    }
}

const UploadToBrowser = ({ children, onChange }) => {
    /*
    const [height, setHeight] = useState(0);
    const [width, setWidth] = useState(0);
    */

    const handleChange = ({target}) => {
        const { files } = target;
        if (files.length < 1) {
            onChange(undefined);
            return;
        }
        const reader = new FileReader();
        reader.onload = ({target}) => {
            const { result } = target;
            onChange(result);
        };
        reader.readAsDataURL(files[0]);
    };

    return (
        <div>
            <input type="file" onChange={handleChange} />
            <Wrapped>{children}</Wrapped>
        </div>
    );
};

export default UploadToBrowser;
