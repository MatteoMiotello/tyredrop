import {useSelector} from "react-redux";
import {selectUser} from "../store/auth-selector";
import {useState} from "react";

const useAuth = () => {
    const user = useSelector( selectUser );
    const [ isLogged, setIsLogged ] = useState( false );
    const [ role, setRole ] = useState( null );


};