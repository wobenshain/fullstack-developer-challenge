import React, { useContext } from 'react';
import {Link} from 'react-router-dom';
import {Icon, notification} from 'antd';
import { serverDown, serverError } from 'Consts/errorCodes';
import ApiContext from 'Contexts/api';

import './index.css';

const Item = ({id, name, description, picture}) => {
    const { itemIdDelete } = useContext(ApiContext);
    function handleClick(e) {
        e.preventDefault();
        e.stopPropagation();
        itemIdDelete(id, (err) => {
            if (!err) {
                window.location.reload();
                return null;
            }
            const {response: {body: { message = serverDown} = {}} = {}} = err;
            notification.error({ message: serverError(message) });
        });
        return false;
    }

    return (
        <Link className="item" to={`/item/${id}`}>
            <section>
                <header>
                    {name}
                    <Icon type="close-circle" onClick={handleClick} />
                </header>
                <div className="img-container">
                    <img src={picture} alt={name} />
                </div>
                <div className="description">{description}</div>
            </section>
        </Link>
    );
};

export default Item;
