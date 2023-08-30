import React from "react";
import {Img} from "react-image";
import {useDispatch} from "react-redux";
import {Link} from "react-router-dom";
import {logout} from "../../modules/auth/store/auth-slice";
import Menu from "../components-library/Menu";
import CartButton from "./CartButton";
import MainLogo from "./Logo";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {
    faBagShopping,
    faFileInvoice,
    faLocationDot,
    faRightFromBracket,
    faUser
} from "@fortawesome/free-solid-svg-icons";
import {useAuth} from "../../modules/auth/hooks/useAuth";

const Navbar: React.FC = () => {
    const auth = useAuth();
    const dispatch = useDispatch();

    return <div className="navbar rounded-box bg-base-100">
        <div className="navbar-start">
            <Link to='/'>
                <MainLogo width={120}/>
            </Link>
            <h2 className="text-xl font-bold uppercase ml-4 font-slogan hidden lg:block text-neutral-500">
                Massimo stock <br/> minimo prezzo
            </h2>
        </div>
        <div className="navbar-center flex gap-4 hidden md:block">
            <Link to="/" className="uppercase btn btn-ghost"> Home </Link>
            <Link to="/contacts" className="uppercase btn btn-ghost"> Contatti </Link>
            <Link to={`/user/${auth.user?.user?.userID}/orders`} className="uppercase btn btn-ghost"> I miei ordini </Link>
        </div>
        <div className="navbar-end">
            {
                auth.user?.user &&
                <div className="mx-4 flex flex-col uppercase font-semibold text-sm">
                    <Link className="link-primary" to={`/user/${auth.user.user.userID}`}> { auth.user.user.email } </Link>
                    <span> Codice utente: #{ String( auth.user?.user?.userCode ).padStart( 5, '0' ) } </span>
                </div>
            }
            <CartButton/>
            <div className="flex-none gap-2">
                <div className="dropdown dropdown-end z-50">
                    <Menu.Dropdown label={<div className="avatar ring-base-300 ring-2 rounded-full placeholder h-full p-1">
                        <div className="text-neutral-content rounded-full" >
                            {
                               auth.user?.user?.avatarUrl ?
                                   <Img className="bg-base-100 !object-contain" src={auth.user.user.avatarUrl}/> :
                                   <span className="text-xl bg-neutral-focus"> <FontAwesomeIcon className="w-36" icon={faUser}/> </span>
                            }
                        </div>
                    </div>}>
                        <Menu.Item>
                            <Link to={`/user/${auth.user?.user?.userID}`}>
                                <FontAwesomeIcon icon={faUser}/> Principale
                            </Link>
                        </Menu.Item>
                        <Menu.Item>
                            <Link to={`/user/${auth.user?.user?.userID}/address`}>
                                <FontAwesomeIcon icon={faLocationDot}/>I miei indirizzi
                            </Link>
                        </Menu.Item>
                        <Menu.Item>
                            <Link to={`/user/${auth.user?.user?.userID}/orders`}>
                                <FontAwesomeIcon icon={faBagShopping}/>I miei ordini
                            </Link>
                        </Menu.Item>
                        <Menu.Item>
                            <Link to={`/user/${auth.user?.user?.userID}/invoices`}>
                                <FontAwesomeIcon icon={faFileInvoice}/>Le mie fatture
                            </Link>
                        </Menu.Item>
                        <Menu.Item>
                            <a onClick={() => dispatch(logout())}>
                                <FontAwesomeIcon icon={faRightFromBracket}/> Esci
                            </a>
                        </Menu.Item>
                    </Menu.Dropdown>
                </div>
            </div>
        </div>
    </div>;
};

export default Navbar;