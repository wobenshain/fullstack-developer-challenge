import React, { useEffect, useState } from 'react';
import {
    BrowserRouter as Router,
    Link,
    Redirect,
    Route,
    Switch,
} from 'react-router-dom';
import {notification} from "antd";
import AddItem from 'Components/AddItem';
import Item from 'Components/Item';
import ItemList from 'Components/ItemList';
import LoadingNotification from "Components/LoadingNotification";
import LoadingContext from 'Contexts/loading';
import { negativeFetch } from 'Consts/errorCodes';
import history from 'Helpers/history';

import 'antd/dist/antd.css';
import './App.css';

const App = () => {
    notification.config({});
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

    return (
        <LoadingContext.Provider value={{ loading, setLoading: setToggle }}>
            <Router history={history}>
                <div className="App">
                    <header className="App-header">
                        <div className="header-limit">
                            <Switch>
                                <Route exact path="/" render={() => null} />
                                <Route exact path="/item/edit/:id" render={({ match: { params: { id = 0 } } }) => <Link to={`/item/${id}`} className="left">Back</Link>} />
                                <Route render={() => <Link to="/" className="left">Home</Link>} />
                            </Switch>
                            <span>Item Inventory</span>
                            <Switch>
                                <Route exact path="/item/edit*" render={() => null} />
                                <Route exact path="/item/:id" render={({ match: { params: { id = 0 } } }) => <Link to={`/item/edit/${id}`} className="right">Edit</Link>} />
                                <Route render={() => <Link to="/item/edit" className="right">Add Item +</Link>} />
                            </Switch>
                        </div>
                    </header>
                    <LoadingNotification />
                    <Switch>
                        <Route exact path="/item/edit" component={AddItem} />
                        <Route exact path="/item/edit/:id" component={AddItem} />
                        <Route exact path="/item/:id" component={Item} />
                        <Route exact path="/" component={ItemList} />
                        <Redirect to="/" />
                    </Switch>
                </div>
            </Router>
        </LoadingContext.Provider>
    );
};

export default App;
