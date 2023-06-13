import React from "react";
import {useAuth} from "../auth/hooks/useAuth";

const UserPage: React.FC = () => {
    const auth = useAuth();

    return <main className="p-4">
        <h3 className="text-xl"> Dati dell'utente </h3>
        <ul>
            <li>Nome: {auth.user?.user?.name}</li>
            <li>Cognome: {auth.user?.user?.surname}</li>
            <li>Email: {auth.user?.user?.email}</li>
        </ul>
    </main>;
};

export default UserPage;