import React, {useEffect, useState} from "react";
import {useLoaderData, useNavigate} from "react-router-dom";
import {FetchUserQuery, User, UserBilling} from "../../__generated__/graphql";
import Panel from "../../common/components-library/Panel";
import Spinner from "../../common/components/Spinner";
import UserBillingPanel from "./components/UserBillingPanel";
import UserEducumentPanel from "./components/UserEducumentPanel";

const UserPage: React.FC = () => {
    const res = useLoaderData() as { data: FetchUserQuery, loading: boolean };
    const [user, setUser] = useState<User | null>(null);
    const [isLoading, setLoading] = useState<boolean>(false);
    const navigate = useNavigate();

    useEffect(() => {
        if (res?.data) {
            setUser(res.data.user as User);
        }

        setLoading(res.loading);
    }, [res]);

    if (isLoading || !user) {
        return <main className="relative">
            <Spinner/>
        </main>;
    }

    return <main className="flex flex-wrap gap-2 w-full">
        <Panel className="flex-auto w-1/2">
            <Panel.Title>
                Dati dell'utente
            </Panel.Title>
            <ul>
                <li><strong>Nome: </strong> {user.name}</li>
                <li><strong>Cognome: </strong> {user.surname}</li>
                <li><strong>Email: </strong> {user.email}</li>
            </ul>
        </Panel>
        {
            user.userBilling &&
            <>
                <UserBillingPanel
                    className="flex-auto w-1/2"
                    userBilling={user.userBilling as UserBilling}
                    onUpdate={ () => navigate('.', {replace: true}) }
                />
                <UserEducumentPanel
                    className="flex-auto"
                    userBilling={user.userBilling as UserBilling}
                    onUpdate={ () => navigate('.', {replace: true}) }
                />
            </>
        }
    </main>;
};

export default UserPage;