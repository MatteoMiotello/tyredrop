import React from "react";
import RegisterForm from "./components/RegisterForm";

const RegisterPage: React.FC = () => {
    return <>
        <div className="flex flex-col justify-center items-center my-auto">
            <RegisterForm/>
        </div>
    </>;
};

export default RegisterPage;