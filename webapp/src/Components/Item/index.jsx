import React, {useContext, useEffect, useState} from 'react';
import ApiContext from 'Contexts/api';
import LoadingContext from 'Contexts/loading';
import { serverDown, serverError } from 'Consts/errorCodes';

import './index.css';
import {notification} from "antd";

const Item = ({ match: { params: { id = 0 } } }) => {
    const [item, setItem] = useState(null);
    const { itemIdGet } = useContext(ApiContext);
    const { setLoading } = useContext(LoadingContext);

    useEffect(() => {
        setLoading(true);
        itemIdGet(id, (err, data) => {
            if (!err) {
                setItem(data);
            }
            setLoading(false);
            if (!err) return null;
            console.log(err.response);
            const {response: { text } = {}} = err;
            const { message = serverDown } = JSON.parse(text);
            notification.error({ message: serverError(message) });
        });
    }, [id]);

    if (!item) return null;
    const { name, description, picture, keyValueFields } = item;
    return (
        <div className="item-display">
            <header>{name}</header>
            <div className="img-container">
                <img src={picture} alt={name} />
            </div>
            <div className="properties">
                <div className="description">{description}</div>
                <ul>
                    { keyValueFields.map(({ id, name, value}) => (
                        <li key={id}>
                            {
                                value
                                    ? <span><strong>{name}:</strong> {value}</span>
                                    : <span><strong>{name}</strong></span>
                            }
                        </li>
                    )) }
                </ul>
            </div>
        </div>
    );
};

export default Item;
