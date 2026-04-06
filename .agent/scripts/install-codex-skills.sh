#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"
SRC_DIR="${1:-${REPO_ROOT}/.agent/skills}"
DEST_DIR="${2:-${REPO_ROOT}/.codex/skills}"

if [[ ! -d "${SRC_DIR}" ]]; then
  echo "Source skill directory not found: ${SRC_DIR}" >&2
  exit 1
fi

mkdir -p "${DEST_DIR}"

linked=0

for skill_dir in "${SRC_DIR}"/*; do
  [[ -d "${skill_dir}" ]] || continue
  [[ -f "${skill_dir}/SKILL.md" ]] || continue

  skill_name="$(basename "${skill_dir}")"
  target="${DEST_DIR}/${skill_name}"

  ln -sfn "${skill_dir}" "${target}"
  echo "linked: ${target} -> ${skill_dir}"
  linked=$((linked + 1))
done

echo
echo "Done. Linked ${linked} skills into ${DEST_DIR}."
echo "Restart Codex to pick up new skills."
