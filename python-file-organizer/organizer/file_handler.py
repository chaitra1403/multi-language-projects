import os
import shutil

# File type mapping
FOLDERS = {
    "images": [".jpg", ".png", ".jpeg"],
    "documents": [".pdf", ".txt", ".docx"],
    "code": [".py", ".java", ".cpp"],
    "archives": [".zip", ".rar"]
}

def organize_folder(path):
    print("Indide organise folder")
    if not os.path.exists(path):
        print("Path does not exist!")
        return

    for file in os.listdir(path):
        file_path = os.path.join(path, file)

        if os.path.isfile(file_path):
            moved = False

            for folder, extensions in FOLDERS.items():
                if any(file.endswith(ext) for ext in extensions):
                    target_folder = os.path.join(path, folder)

                    os.makedirs(target_folder, exist_ok=True)
                    shutil.move(file_path, os.path.join(target_folder, file))

                    print(f"Moved {file} → {folder}/")
                    moved = True
                    break

            if not moved:
                other_folder = os.path.join(path, "others")
                os.makedirs(other_folder, exist_ok=True)
                shutil.move(file_path, os.path.join(other_folder, file))
