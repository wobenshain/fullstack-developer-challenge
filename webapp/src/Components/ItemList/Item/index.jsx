import React from 'react';
import {Link} from 'react-router-dom';

import './index.css';

const Item = ({id, name, description, picture, keyValueFields}) => (
    <Link className="item" to={`/item/${id}`}>
        <section>
            <header>{name}</header>
            <div className="img-container">
                <img src={picture} alt={name} />
            </div>
            <div className="description">{description}</div>
        </section>
    </Link>
);

export default Item;
