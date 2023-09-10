import React from "react";
import {Link, useRouteError} from "react-router-dom";

const ErrorPage: React.FC = () => {
    const err = useRouteError() as any;

    return <main className="w-screen h-screen flex flex-col items-center justify-center bg-base-300 space-y-10">
        <h1 className="text-4xl">

        </h1>
        <h2 className="font-bold text-5xl">
            {err.status ?? 404}
        </h2>
        <Link className="btn btn-lg btn-primary" to="/">
            Torna alla Home
        </Link>
    </main>;
};

export default ErrorPage;