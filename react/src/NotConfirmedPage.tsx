import {useTranslation} from "react-i18next";
import Logo from "./common/components/Logo";
import CustomFooter from "./common/components/CustomFooter";
import {useAuth} from "./modules/auth/hooks/useAuth";
import {useNavigate} from "react-router-dom";
import {useEffect} from "react";
import Button from "./common/components-library/Button";

const NotConfirmedPage: React.FC = () => {
    const {t} = useTranslation();
    const navigate = useNavigate();
    const auth = useAuth();


    useEffect(() => {
        if (auth.isUserCompleted()) {
            navigate('/');

            return;
        }
    }, [auth]);

    return <>
        <main className="flex flex-col justify-center items-center h-screen">
            <div className="bg-base-100 rounded-box flex flex-col items-center text-center p-24">
                <Logo width={150}/>
                <h1 className="my-10 text-2xl font-medium"> {t('not_confirmed.title')} </h1>
                <p> {t('not_confirmed.text')} </p>
                <Button className="mt-10" type="primary" onClick={() => navigate( '/' )}> Home </Button>
            </div>
        </main>
        <CustomFooter/>
    </>;
};

export default NotConfirmedPage;