import os
from pathlib import Path

def rename_images_by_type(base_path):
    # 機械タイプごとのカウンターと変換マップ
    type_mapping = {
        'lathe_machine': 'lathe',
        'press_machine': 'press',
        'radial_drilling_machine': 'drill',
        'milling_machine': 'mill'
    }
    
    counters = {folder: 1 for folder in type_mapping.keys()}
    
    # 各機械タイプのフォルダをループ
    for folder, prefix in type_mapping.items():
        folder_path = Path(base_path) / folder
        
        # フォルダが存在する場合のみ処理
        if folder_path.exists():
            # jpgファイルを取得してソート
            jpg_files = sorted([f for f in folder_path.glob('*.jpg')])
            
            # 各ファイルの名前を連番で変更
            for jpg_file in jpg_files:
                new_name = f"{prefix}{counters[folder]}.jpg"
                new_path = folder_path / new_name
                
                # ファイル名変更
                jpg_file.rename(new_path)
                print(f"Renamed: {jpg_file.name} -> {new_name}")
                
                # カウンターをインクリメント
                counters[folder] += 1

if __name__ == "__main__":
    # スクリプトの現在のディレクトリを取得
    current_path = os.path.dirname(os.path.abspath(__file__))
    # machineryフォルダへの正しいパスを構築
    machinery_path = os.path.join(current_path, "..", "src", "assets", "machinery")
    # 絶対パスに変換
    machinery_path = os.path.abspath(machinery_path)
    
    # パスが存在することを確認
    if not os.path.exists(machinery_path):
        print(f"エラー: ディレクトリが見つかりません: {machinery_path}")
        exit(1)
        
    rename_images_by_type(machinery_path)