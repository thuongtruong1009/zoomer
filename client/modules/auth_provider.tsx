import { useState, createContext, useEffect } from "react";
import { useRouter } from "next/router";

export type UserInfo = {
  username: string;
  userId: string;
  token: string;
};

export const AuthContext = createContext<{
  authenticated: boolean;
  setAuthenticated: (auth: boolean) => void;
  user: UserInfo;
  setUser: (user: UserInfo) => void;
}>({
  authenticated: false,
  setAuthenticated: () => {},
  user: { username: "", userId: "", token: "" },
  setUser: () => {},
});

const AuthContextProvider = ({ children }: { children: React.ReactNode }) => {
  const [authenticated, setAuthenticated] = useState(false);
  const [user, setUser] = useState<UserInfo>({
    username: "",
    userId: "",
    token: "",
  });

  const router = useRouter();

  useEffect(() => {
    const userInfo = localStorage.getItem("user_info");

    if (!userInfo) {
      if (window.location.pathname != "/signup") {
        router.push("/login");
        return;
      }
    } else {
      const user: UserInfo = JSON.parse(userInfo);
      if (user) {
        setUser({
          username: user.username,
          userId: user.userId,
          token: user.token,
        });
      }
      setAuthenticated(true);
    }
  }, [authenticated]);

  return (
    <AuthContext.Provider
      value={{
        authenticated: authenticated,
        setAuthenticated: setAuthenticated,
        user: user,
        setUser: setUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContextProvider;
