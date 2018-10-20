/**
 * Created by cahyo on 01/22/2018.
 */

export const timeout = (ms: number) => {
    return new Promise((res) => setTimeout(res, ms));
  };
  