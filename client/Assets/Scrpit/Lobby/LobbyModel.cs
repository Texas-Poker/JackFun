namespace JackFun.Lobby
{
    public class GameInfo
    {
        public uint GameId { get; set; }
        public string GameName { get; set; }
        public bool IsOpen { get; set; }

        public GameInfo(Pb.Lobby.RespLobbyInfo.Types.LobbyInfo info)
        {
            GameId = info.GameId;
            GameName = info.GameName;
            IsOpen = info.IsOpen;
        }
    }
}