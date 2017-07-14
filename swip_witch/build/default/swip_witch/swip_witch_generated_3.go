components {
  id: "script"
  component: "/swip_witch/swip_witch.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "coinfactory"
  type: "factory"
  data: "prototype: \"/swip_witch/witch.go\"\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "coinmagnetcollisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_KINEMATIC\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"coinmagnet\"\nmask: \"coin\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 1\n  }\n  data: 200.0\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_KINEMATIC\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"player\"\nmask: \"coin\"\nmask: \"powerup\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 1\n  }\n  data: 40.0\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "powerupfactory"
  type: "factory"
  data: "prototype: \"/swip_witch/powerup.go\"\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/swip_witch/swip_witch.atlas\"\ndefault_animation: \"littlewitch\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
