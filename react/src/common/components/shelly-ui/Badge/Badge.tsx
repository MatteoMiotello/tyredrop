import clsx from "clsx";
import React, {HTMLAttributes, PropsWithChildren, useEffect, useState} from "react";
import {twMerge} from "tailwind-merge";
import {calculateLuminance} from "../../../utilities/text-color";
import {swtc} from "../utils";

type BadgeProps = {
    badgeType?: 'neutral' | 'primary' | 'secondary' | 'accent' | 'ghost' | undefined,
    color?: string,
    outline?: boolean
} & HTMLAttributes<HTMLDivElement> & PropsWithChildren

const Badge: React.FC<BadgeProps> = ({children, className, badgeType, color, outline, ...props}) => {
    const [textClass, setTextClass] = useState( 'text-black' );

    useEffect(() => {
        if ( color ) {
            setTextClass( (calculateLuminance( color ) > 0.5) ? 'text-black' : 'text-white' );
        }
    }, [color]);

    const classNames = twMerge(
        'badge',
        'h-auto',
        clsx(
            badgeType && swtc(badgeType, {
                neutral: 'badge-neutral',
                primary: 'badge-primary',
                secondary: 'badge-secondary',
                accent: 'badge-accent',
                ghost: 'badge-ghost',
            }),
            outline && 'badge-outline',
        ),
        textClass,
        className,
    );

    return <div className={classNames}
                style={{backgroundColor: (!outline) ? color : undefined, borderColor: color}} {...props}>
        {children}
    </div>;
};

export default Badge;