import * as dotenv from 'dotenv'

dotenv.config()

const APEnv = ["ROOT_TOKEN", "BASEURL", "READ_TOKEN", "WRITE_TOKEN"] as const;

type APEnvType = typeof APEnv[number];

type APEnvMapType = { [key in APEnvType]: string };

const getEnv = (key: APEnvType): string => {
    const result = process.env[key];
    if (result === undefined) {
        throw new Error(`${key} not found`);
    } else {
        return result;
    }
};

export const env = APEnv.map((value) => ({ [value]: getEnv(value) })).reduce(
    (acc, cur) => ({ ...acc, ...cur }),
    {}
) as APEnvMapType;
