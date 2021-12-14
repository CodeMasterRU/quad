<?php
function findIn(string $needle, array $tab): bool
{
    foreach ($tab as $key => $value) {
        if ($key == $needle) {
            return $value;
        } else if (is_array($value)) {
            $key_result = findIn($needle, $value);
            if ($key_result !== false) {
                return $key_result;
            }
        }
    }
    return false;
}