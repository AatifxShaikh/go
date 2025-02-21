Here’s a **cheat sheet** of commonly used **Neovim key mappings** to boost your workflow:  

---

## **BASICS**  
| Action                     | Command |
|----------------------------|---------|
| **Save file**             | `:w` |
| **Save and quit**         | `:wq` or `ZZ` |
| **Quit without saving**   | `:q!` |
| **Close current buffer**  | `:bd` |
| **Undo**                  | `u` |
| **Redo**                  | `Ctrl + r` |
| **Delete a line**         | `dd` |
| **Copy (yank) a line**    | `yy` |
| **Paste below cursor**    | `p` |
| **Paste above cursor**    | `P` |

---

## **NAVIGATION**  
| Action                     | Command |
|----------------------------|---------|
| **Move to start of line**  | `0` |
| **Move to end of line**    | `$` |
| **Jump to beginning of file** | `gg` |
| **Jump to end of file**        | `G` |
| **Move up a paragraph**    | `{` |
| **Move down a paragraph**  | `}` |
| **Move by words forward**  | `w` |
| **Move by words backward** | `b` |
| **Search for text**        | `/text` then `n` (next), `N` (previous) |

---

## **WINDOWS (SPLITS)**  
| Action                    | Command |
|---------------------------|---------|
| **Split window horizontally**  | `:sp` |
| **Split window vertically**    | `:vs` |
| **Switch between splits**      | `Ctrl + w + w` |
| **Resize split (increase width)** | `Ctrl + w >` |
| **Resize split (decrease width)** | `Ctrl + w <` |
| **Resize split (increase height)** | `Ctrl + w +` |
| **Resize split (decrease height)** | `Ctrl + w -` |

---

## **TABS**  
| Action                    | Command |
|---------------------------|---------|
| **Open new tab**          | `:tabnew` |
| **Next tab**              | `gt` |
| **Previous tab**          | `gT` |
| **Close current tab**     | `:tabclose` |

---

## **BUFFERS**  
| Action                    | Command |
|---------------------------|---------|
| **List buffers**          | `:ls` |
| **Switch buffer**         | `:b buffer_number` |
| **Next buffer**           | `:bn` |
| **Previous buffer**       | `:bp` |
| **Delete buffer**         | `:bd` |

---

## **TEXT EDITING**  
| Action                    | Command |
|---------------------------|---------|
| **Indent right**          | `>>` |
| **Indent left**           | `<<` |
| **Replace a character**   | `r<char>` |
| **Replace entire word**   | `ciw` |
| **Delete word**           | `dw` |
| **Change word**           | `cw` |
| **Select entire file**    | `ggVG` |
| **Toggle case (uppercase/lowercase)** | `~` |

---

## **VISUAL MODE**  
| Action                    | Command |
|---------------------------|---------|
| **Enter visual mode**     | `v` |
| **Enter line visual mode** | `V` |
| **Enter block visual mode** | `Ctrl + v` |
| **Copy (yank) selection** | `y` |
| **Cut (delete) selection** | `d` |
| **Paste selection**       | `p` |

---

## **PLUGIN-SPECIFIC (If Using LSP, Treesitter, Telescope, etc.)**  

#### **LSP (Language Server Protocol)**
| Action                         | Command |
|---------------------------------|---------|
| **Show LSP hover (docs)**      | `K` |
| **Go to definition**           | `gd` |
| **Go to references**           | `gr` |
| **Show errors**                | `:TroubleToggle` or `:LspDiagnostics` |
| **Format code**                | `:lua vim.lsp.buf.format()` |

#### **Telescope (Fuzzy Finder)**
| Action                           | Command |
|-----------------------------------|---------|
| **Find files**                   | `:Telescope find_files` |
| **Find text in files**           | `:Telescope live_grep` |
| **Search buffers**               | `:Telescope buffers` |
| **Search help tags**             | `:Telescope help_tags` |

---

This cheat sheet should cover most of your **Neovim essentials**! Let me know if you need any specific mappings. 🚀
