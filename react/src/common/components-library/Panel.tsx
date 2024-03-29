import React, {PropsWithChildren} from "react";

interface PanelProps extends PropsWithChildren {
    className?: string
}

const Panel : React.FC<PanelProps> = ( props: PanelProps ) => {
    return <div className={"rounded-box shadow bg-base-100 p-4 relative " + props.className}>
        { props.children }
    </div>;
};

type TitleProps = PropsWithChildren
const Title: React.FC<TitleProps> = ({children} ) => {
    return <h3 className="text-xl font-semibold my-2 flex justify-between items-center">
        {children}
    </h3>;
};

export default Object.assign( Panel, {
    Title
} );