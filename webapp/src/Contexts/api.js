import { createContext } from 'react';
import DefaultApi from 'Swagger/src/api/DefaultApi';

export default createContext(new DefaultApi());
