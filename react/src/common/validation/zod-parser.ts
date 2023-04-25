import { ZodError, ZodType} from "zod";

export const zodParser = (type: ZodType, value: any): string | null => {
    try {
        type.parse(value);
    } catch (e: any) {
        if (e instanceof ZodError) {
            if (!e.isEmpty) {
                return e.errors[0].message;
            }
        }

        if (e instanceof Error) {
            return e.message;
        }

        if (typeof e === 'string') {
            return e;
        }
    }

    return null;
};