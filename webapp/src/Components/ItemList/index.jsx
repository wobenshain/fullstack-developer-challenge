import React, {useContext, useEffect, useState} from 'react';
import { notification } from 'antd';
import Item from 'Components/ItemList/Item';
import ApiContext from 'Contexts/api';
import LoadingContext from 'Contexts/loading';
import { serverDown, serverError } from 'Consts/errorCodes';

import './index.css';

const ItemList = () => {
    const [items, setItems] = useState([]);
    const { itemGet } = useContext(ApiContext);
    const { setLoading } = useContext(LoadingContext);

    useEffect(() => {
        setLoading(true);
        itemGet((err, data) => {
            if (!err) {
                setItems(data);
            }
            setLoading(false);
            if (!err) return null;
            const {response: {body: { message = serverDown} = {}} = {}} = err;
            notification.error({ message: serverError(message) });
        });
    }, [itemGet]);

    return (
        <div className="item-list">
            {
                items.map(({id, ...remainder}) => (
                    <Item key={id} id={id} {...remainder} />
                ))
            }
        </div>
    )
};

export default ItemList;
