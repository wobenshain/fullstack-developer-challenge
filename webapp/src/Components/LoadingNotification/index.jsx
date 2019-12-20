import React, { useContext } from 'react';
import { Spin } from 'antd';
import LoadingContext from 'Contexts/loading';

const LoadingNotification = () => {
    const { loading } = useContext(LoadingContext);
    const display = loading ? 'block' : 'none';
    return (
        <div style={{ display }}>
            Loading...
            <Spin />
        </div>
    );
};

export default LoadingNotification;
