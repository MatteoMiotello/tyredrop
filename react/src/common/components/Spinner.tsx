// @ts-ignore
import {ReactComponent as TyreSpinner} from '../../assets/car-tire-wheel-icon.svg' ;

const Spinner: React.FC = () => {
    return <div className="absolute flex justify-center items-center z-50 w-full h-full top-50 bg-transparent/20 bg-grey-50">
        <TyreSpinner className="animate-spin animate" style={{fill: "#fff", width:"50px"}}></TyreSpinner>
    </div>;
};

export default Spinner;