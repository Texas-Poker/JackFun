using System;

namespace JackFun
{
    public class Session
    {
        public static string Token { get; set; }

        public static User User { get; set; }
    }


    public class User
    {
        public string NickName { get; set; }
        public Pb.Enum.Sex Sex { get; set; }
        public UInt32 Gold { get; set; }
        public UInt32 Diamond { get; set; }
        public UInt32 Lv { get; set; }
        public long UID { get; set; }

        public User(Pb.Lobby.RespAuth info)
        {
            UID = info.UID;
            NickName = info.NickName;
            Sex = info.Sex;
            Gold = info.Gold;
            Diamond = info.Diamond;
            Lv = info.Lv;
        }
    }
}