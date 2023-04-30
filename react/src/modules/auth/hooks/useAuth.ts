import {useSelector} from "react-redux";
import {selectAuthStatus} from "../store/auth-selector";

export const useAuth = () => {
    return useSelector(selectAuthStatus);
};