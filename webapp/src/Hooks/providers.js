import {useEffect, useState} from "react";
import {notification} from "antd";
import {negativeFetch} from "../Consts/errorCodes";

export function useLoading() {
    const [loading, setLoading] = useState(false)
    const [toggle, setToggle] = useState(null);
    const [outstanding, setOutstanding] = useState(0);

    useEffect(() => {
        if (toggle !== null) {
            if (toggle) {
                setOutstanding(outstanding + 1);
            } else {
                let newValue = outstanding - 1;
                if (newValue < 0) {
                    newValue = 0;
                    notification.warn({ message: negativeFetch });
                }
                setOutstanding(newValue);
            }
            setToggle(null);
        }
    }, [toggle]);

    useEffect(() => {
        setLoading(outstanding > 0);
    }, [outstanding]);

    return { loading, setLoading: setToggle };
}
