import React from "react";
import {useDispatch} from "react-redux";
import {LoginRequest} from "../../common/backend/requests/LoginRequest";
import Panel from "../../common/components-library/Panel";
import logo from "../../assets/logo-transparent.png";
import LoginForm from "./components/LoginForm";

export const LoginPage: React.FC = () => {
    const dispatch = useDispatch();

    const login = ( loginRequest: LoginRequest ) => {
        dispatch({ type: 'auth/login', payload: loginRequest });
    };

    return <div>
        <main className="bg-base lg:p-24 p-4">
            <Panel className="flex flex-col justify-center items-center my-auto">
                <img src={logo} width={100} alt={"Logo"}/>
                <LoginForm login={ login }/>
            </Panel>
        </main>
    </div>;
};