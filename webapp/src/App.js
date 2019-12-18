import React from 'react';
import {
    BrowserRouter as Router,
    Link,
    Redirect,
    Route,
    Switch,
} from 'react-router-dom';
import AddItem from './Components/AddItem';
import Item from './Components/Item';
import ItemList from './Components/ItemList';

import './App.css';

const App = () => {
    return (
        <Router>
            <div className="App">
                <header className="App-header">
                    <Switch>
                        <Route exact path="/" render={() => null} />
                        <Route exact path="/item/edit/:id" render={({ match: { params: { id = 0 } } }) => <Link to={`/item/${id}`}>Back</Link>} />
                        <Route render={() => <Link to="/">Home</Link>} />
                    </Switch>
                    <span>Item Inventory</span>
                    <Switch>
                        <Route exact path="/item/edit" render={() => null} />
                        <Route render={() => <Link to="/item/edit">Add Item +</Link>} />
                    </Switch>
                </header>
                <Switch>
                    <Route exact path="/item/edit" component={AddItem} />
                    <Route exact path="/item/edit/:id" component={AddItem} />
                    <Route exact path="/item/:id" component={Item} />
                    <Route exact path="/" component={ItemList} />
                    <Redirect to="/" />
                </Switch>
            </div>
        </Router>
    );
};

export default App;
