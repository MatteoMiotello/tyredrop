import React, {Suspense} from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import Spinner from "./common/components/Spinner";
import Root from "./Root";


ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <Suspense fallback={<Spinner/>}>
        <React.StrictMode>
            <Root/>
        </React.StrictMode>
    </Suspense>
);
