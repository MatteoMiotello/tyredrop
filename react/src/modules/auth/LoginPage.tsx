import React from "react";
import Panel from "../../common/components-library/Panel";
import logo from "../../assets/logo-transparent.png";
import useLogin from "../../hooks/backend/useLogin";
import LoginForm, {LoginRequest} from "./components/LoginForm";

export const LoginPage: React.FC = () => {
    const [data, error, handleLogin] = useLogin();

    const login = ( req: LoginRequest ) => {

    };

    return <div>
        <main className="bg-base lg:p-24 p-4">
            <Panel className="flex flex-col justify-center items-center my-auto">
                <img src={logo} width={100} alt={"Logo"}/>
                <LoginForm login={login}/>
            </Panel>
        </main>
    </div>;
};