using System.Collections.Generic;
using FairyGUI;
using UnityEngine;

namespace JackFun.UI
{
    public static class UIUtil
    {
        private static Dictionary<string, UIPanel> _dic = new Dictionary<string, UIPanel>();

        public static void Init()
        {
            //common 的package
            UIPackage.AddPackage("UI/FP_Common");
            UIPackage.AddPackage("UI/FP_Common_Head_Item");

            //common 的UI
            OpenUI<UILoading>();
            OpenUI<UITips>();
        }

        public static void OpenUI<T>() where T : UIPanel
        {
            var go = new GameObject
            {
                layer = LayerMask.NameToLayer("UI"),
            };
            var uiPanel = go.AddComponent<T>();
            var typeT = typeof(T);
            var prefabName = "Panel_" + typeT.Name;
            go.name = prefabName;
            UIPackage.AddPackage("UI/" + uiPanel.packageName);
            uiPanel.fitScreen = FitScreen.FitHeightAndSetCenter;
            uiPanel.container.renderMode = RenderMode.ScreenSpaceOverlay;
            uiPanel.ui.MakeFullScreen();
            uiPanel.CreateUI();
            _dic[prefabName] = uiPanel;
        }

        public static void CloseUI<T>() where T : UIPanel
        {
            var typeT = typeof(T);
            var prefabName = "Panel_" + typeT.Name;

            if (!_dic.ContainsKey(prefabName))
            {
                return;
            }

            var target = _dic[prefabName];
            _dic.Remove(prefabName);
            target.container.Dispose();
            target.ui.Dispose();
            Object.Destroy(target.gameObject);
        }
    }
}