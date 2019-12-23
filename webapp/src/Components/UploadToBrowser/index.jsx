import React from 'react';

import './index.css';

const UploadToBrowser = ({ children, className = '', onChange }) => {
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
        <div className={`UploadToBrowser ${className}`.trim()}>
            {children}
            <input type="file" onChange={handleChange} />
        </div>
    );
};

export default UploadToBrowser;
