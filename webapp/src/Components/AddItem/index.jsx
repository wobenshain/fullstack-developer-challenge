import React, { useContext, useEffect, useState } from 'react';
import {Button, Icon, Input, notification} from 'antd';
import Upload from "Components/UploadToBrowser";
import ApiContext from 'Contexts/api';
import LoadingContext from 'Contexts/loading';
import { serverDown, serverError } from 'Consts/errorCodes';

import KeyValueList from "./KeyValueList";

const { TextArea } = Input;

const onChange = (set) => ({ target }) => {
    const { value } = target;
    set(value);
};

export const AddItem = ({ history, match: { params: { id = 0 } } }) => {
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [picture, setPicture] = useState(undefined);
    const [keyValueFields, setKeyValueFields] = useState([{ id: 1, name: "", value: ""}]);
    const { itemIdGet, itemIdPatch, itemPost } = useContext(ApiContext);
    const { setLoading } = useContext(LoadingContext);
    useEffect(() => {
        if (id > 0) {
            setLoading(true);
            itemIdGet(id, (err, data) => {
                if (!err) {
                    const { name, description, picture, keyValueFields = [] } = data;
                    setName(name);
                    setDescription(description);
                    setPicture(picture);
                    const kvId = keyValueFields.map(({ id }) => id).reduce((a, b) => Math.max(a, b), 0) + 1;
                    keyValueFields.push({ id: kvId, name: "", value: "" });
                    setKeyValueFields(keyValueFields);
                }
                setLoading(false);
                if (!err) return null;
                const {response: {body: { message = serverDown} = {}} = {}} = err;
                notification.error({ message: serverError(message) });
            });
        }
    }, [id]);

    function onUpsert(err, { id = 0 }) {
        setLoading(false);
        if (!err) {
            if (id) {
                history.push(`/item/${id}`);
            } else {
                history.push('/item');
            }
        }
        if (!err) return null;
        const {response: {body: { message = serverDown} = {}} = {}} = err;
        notification.error({ message: serverError(message) });
    }

    function handle(id, name, value) {
        const nvPair = { name, value };
        let kvList = keyValueFields.slice().map(kv => kv.id === id ? Object.assign(kv, nvPair) : kv);
        kvList = kvList.filter(({name, value}) => name || value);
        const kvId = keyValueFields.map(({ id }) => id).reduce((a, b) => Math.max(a, b), 0) + 1;
        kvList.push({ id: kvId, name: "", value: ""});
        setKeyValueFields(kvList);
    }

    function handleClick() {
        const kvList = keyValueFields.slice(0, keyValueFields.length - 1);
        const request = { item: {
            name, description, picture, keyValueFields: kvList,
        } };
        setLoading(true);
        if (id) {
            request.item.id = Number(id);
            itemIdPatch(id, request, onUpsert)
        } else {
            itemPost(request, onUpsert)
        }
    }

    return (
        <div>
            <Input onChange={onChange(setName)} value={name} />
            <Upload onChange={setPicture}>
                {
                    picture
                        ? <img src={picture} alt={name} style={{ width: '100%' }} />
                        : (
                            <div>
                                <Icon type="plus" />
                                <div className="ant-upload-text">Upload</div>
                            </div>
                        )}
            </Upload>
            <TextArea onChange={onChange(setDescription)} value={description} />
            <KeyValueList kvList={keyValueFields} onChange={handle} />
            <Button onClick={handleClick}>{ id ? 'Update' : 'Create' }</Button>
        </div>
    );
};

export default AddItem;
