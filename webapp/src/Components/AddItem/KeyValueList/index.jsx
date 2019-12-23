import React, { useEffect, useState } from 'react';
import { Input } from 'antd';
import { setStateOnChange } from 'Helpers/events';

import './index.css';

const KeyValuePair = ({ onChange, id, name, value }) => {
    const [localName, setName] = useState(name);
    const [localValue, setValue] = useState(value);
    useEffect(() => { setName(name); }, [name]);
    useEffect(() => { setValue(value); }, [value]);

    const handleBlur = () => {
        onChange(id, localName, localValue);
    };

    return (
        <li>
            <Input onChange={setStateOnChange(setName)} onBlur={handleBlur} placeholder="Property Name" value={localName} />
            <Input onChange={setStateOnChange(setValue)} onBlur={handleBlur} placeholder="Property Value" value={localValue} />
        </li>
    );
};

const KeyValueList = ({ kvList, onChange }) => {
    const handleChange = (id, name, value) => {
        onChange(id, name, value);
    };
    return (
        <ul className="KeyValueList">
            {
                kvList.map(({id, ...kv}) => (
                    <KeyValuePair key={id} id={id} {...kv} onChange={handleChange} />
                ))
            }
        </ul>
    );
};

export default KeyValueList;
