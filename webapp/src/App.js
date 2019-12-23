import React from 'react';
import {
    BrowserRouter as Router,
    Link,
    Redirect,
    Route,
    Switch,
} from 'react-router-dom';
import AddItem from 'Components/AddItem';
import Item from 'Components/Item';
import ItemList from 'Components/ItemList';
import LoadingNotification from "Components/LoadingNotification";
import LoadingContext from 'Contexts/loading';
import history from 'Helpers/history';
import { useLoading } from "Hooks/providers";

import 'antd/dist/antd.css';
import './App.css';

const App = () => {
    return (
        <LoadingContext.Provider value={useLoading()}>
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
